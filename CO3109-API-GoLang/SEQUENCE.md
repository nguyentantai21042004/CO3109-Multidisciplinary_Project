# Authentication API Flow

## 1. Register API Flow
```mermaid
sequenceDiagram
    participant Client
    participant API
    participant UseCase
    participant UserRepo
    participant Encrypt

    Client->>API: POST /api/v1/auth/register
    Note over Client,API: Body: {email, password}
    
    API->>UseCase: Register(ctx, scope, input)
    
    %% Verify Phase
    rect rgb(255, 240, 240)
        UseCase->>UserRepo: GetOne(ctx, scope, {email})
        alt User exists
            UserRepo-->>UseCase: Error (Email existed)
            UseCase-->>API: Error (Email existed)
            API-->>Client: 400 Bad Request
        end
    end
    
    %% Authenticate Phase
    rect rgb(240, 255, 240)
        UseCase->>Encrypt: Encrypt(password)
        Encrypt-->>UseCase: Encrypted password
    end
    
    %% Fetch Phase
    rect rgb(240, 240, 255)
        UseCase->>UserRepo: Create(ctx, scope, {email, encryptedPassword, provider: web, isVerified: false})
        UserRepo-->>UseCase: User created
    end
    
    %% Process Phase
    rect rgb(255, 255, 240)
        UseCase-->>API: Success
        API-->>Client: 200 OK
    end
```

## 2. Send OTP API Flow

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant UseCase
    participant UserRepo
    participant Encrypt
    participant OTP
    participant Email
    participant RabbitMQ

    Client->>API: POST /api/v1/auth/send-otp
    Note over Client,API: Body: {email, password}
    
    API->>UseCase: SendOTP(ctx, scope, input)
    
    %% Verify Phase
    rect rgb(255, 240, 240)
        UseCase->>UserRepo: GetOne(ctx, scope, {email})
        alt User not found
            UserRepo-->>UseCase: Error (User not found)
            UseCase-->>API: Error (User not found)
            API-->>Client: 404 Not Found
        end
        
        UseCase->>Encrypt: Decrypt(user.PasswordHash)
        Encrypt-->>UseCase: Decrypted password
        alt Wrong password
            UseCase-->>API: Error (Wrong password)
            API-->>Client: 400 Bad Request
        end
        
        alt User already verified
            UseCase-->>API: Error (User verified)
            API-->>Client: 400 Bad Request
        end
    end
    
    %% Authenticate Phase
    rect rgb(240, 255, 240)
        UseCase->>UseCase: Check OTP expired
        alt OTP expired or not exists
            UseCase->>OTP: GenerateOTP(now)
            OTP-->>UseCase: {otp, otpExpiredAt}
        end
    end
    
    %% Fetch Phase
    rect rgb(240, 240, 255)
        UseCase->>UserRepo: UpdateVerified(ctx, scope, {userID, otp, otpExpiredAt, isVerified: false})
        UserRepo-->>UseCase: User updated
        
        UseCase->>Email: NewEmail(ctx, {recipient, templateType}, {name, email, otp, otpExpireMin})
        Email-->>UseCase: Email template
    end
    
    %% Process Phase
    rect rgb(255, 255, 240)
        UseCase->>RabbitMQ: PubSendEmailMsg(ctx, scope, {recipient, subject, body})
        UseCase-->>API: Success
        API-->>Client: 200 OK
    end
```

## 3. Verify OTP API Flow

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant UseCase
    participant UserRepo

    Client->>API: POST /api/v1/auth/verify-otp
    Note over Client,API: Body: {email, otp}
    
    API->>UseCase: VerifyOTP(ctx, scope, input)
    
    %% Verify Phase
    rect rgb(255, 240, 240)
        UseCase->>UserRepo: GetOne(ctx, scope, {email})
        alt User not found
            UserRepo-->>UseCase: Error (User not found)
            UseCase-->>API: Error (User not found)
            API-->>Client: 404 Not Found
        end
        
        UseCase->>UseCase: Check OTP match
        alt OTP not match
            UseCase-->>API: Error (Wrong OTP)
            API-->>Client: 400 Bad Request
        end
        
        UseCase->>UseCase: Check OTP expired
        alt OTP expired
            UseCase-->>API: Error (OTP expired)
            API-->>Client: 400 Bad Request
        end
    end
    
    %% Authenticate Phase
    rect rgb(240, 255, 240)
        UseCase->>UserRepo: UpdateVerified(ctx, scope, {userID, isVerified: true})
        UserRepo-->>UseCase: User updated
    end
    
    %% Fetch Phase
    rect rgb(240, 240, 255)
        Note over UseCase: No fetch operations needed
    end
    
    %% Process Phase
    rect rgb(255, 255, 240)
        UseCase-->>API: Success
        API-->>Client: 200 OK
    end
```

## 4. Login API Flow

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant UseCase
    participant UserRepo
    participant Encrypt
    participant Scope
    participant Session
    participant Role

    Client->>API: POST /api/v1/auth/login
    Note over Client,API: Body: {email, password, remember, userAgent, ipAddress, deviceName}
    
    API->>UseCase: Login(ctx, scope, input)
    
    %% Verify Phase
    rect rgb(255, 240, 240)
        UseCase->>UserRepo: GetOne(ctx, scope, {email})
        alt User not found
            UserRepo-->>UseCase: Error (User not found)
            UseCase-->>API: Error (User not found)
            API-->>Client: 404 Not Found
        end
        
        alt User not verified
            UseCase-->>API: Error (User not verified)
            API-->>Client: 400 Bad Request
        end
        
        UseCase->>Encrypt: Decrypt(user.PasswordHash)
        Encrypt-->>UseCase: Decrypted password
        alt Wrong password
            UseCase-->>API: Error (Wrong password)
            API-->>Client: 400 Bad Request
        end
    end
    
    %% Authenticate Phase
    rect rgb(240, 255, 240)
        UseCase->>UseCase: Calculate token expiry
        alt Remember me
            UseCase->>UseCase: Set refresh expiry to 30 days
        else
            UseCase->>UseCase: Set refresh expiry to 7 days
        end
        
        UseCase->>Scope: CreateToken(ctx, {userID, email, type: access})
        Scope-->>UseCase: Access token
    end
    
    %% Fetch Phase
    rect rgb(240, 240, 255)
        UseCase->>UseCase: Generate refresh token (UUID)
        UseCase->>Session: Create(ctx, scope, {userID, accessToken, refreshToken, expiresAt, userAgent, ipAddress, deviceName})
        Session-->>UseCase: Session created
        
        UseCase->>Role: Detail(ctx, scope, user.RoleID)
        Role-->>UseCase: Role details
    end
    
    %% Process Phase
    rect rgb(255, 255, 240)
        UseCase-->>API: Success
        API-->>Client: 200 OK
        Note over Client,API: Response: {user, role, token: {accessToken, refreshToken, expiresAt, sessionID, tokenType: Bearer}}
    end
```

## 5. Social Login API Flow

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant UseCase
    participant OAuth2
    participant UserRepo
    participant Session
    participant Role

    Client->>API: GET /api/v1/auth/social-login/{provider}
    Note over Client,API: Query: {provider: facebook/google/gitlab, redirect: boolean}
    
    API->>UseCase: SocialLogin(ctx, scope, input)
    
    %% Verify Phase
    rect rgb(255, 240, 240)
        UseCase->>UseCase: Validate provider
        alt Invalid provider
            UseCase-->>API: Error (Invalid provider)
            API-->>Client: 400 Bad Request
        end
    end
    
    %% Authenticate Phase
    rect rgb(240, 255, 240)
        UseCase->>OAuth2: GetAuthURL(provider)
        OAuth2-->>UseCase: Authorization URL
    end
    
    %% Process Phase
    rect rgb(255, 255, 240)
        UseCase-->>API: Success (URL)
        alt Redirect true
            API-->>Client: 302 Redirect to OAuth URL
        else
            API-->>Client: 200 OK with URL
        end
    end
```

## 6. Social Callback API Flow

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant UseCase
    participant OAuth2
    participant UserRepo
    participant Session
    participant Role

    Client->>API: GET /api/v1/auth/callback/{provider}
    Note over Client,API: Query: {code, state}
    
    API->>UseCase: SocialCallback(ctx, scope, input)
    
    %% Verify Phase
    rect rgb(255, 240, 240)
        UseCase->>UseCase: Validate provider
        alt Invalid provider
            UseCase-->>API: Error (Invalid provider)
            API-->>Client: 400 Bad Request
        end
    end
    
    %% Authenticate Phase
    rect rgb(240, 255, 240)
        UseCase->>OAuth2: GetUserInfo(provider, code)
        OAuth2-->>UseCase: Social user info
        
        UseCase->>UserRepo: GetOneByEmail(ctx, scope, email)
        alt User not found
            UseCase->>UserRepo: Create(ctx, scope, {email, provider, socialID})
        end
    end
    
    %% Fetch Phase
    rect rgb(240, 240, 255)
        UseCase->>UseCase: Generate tokens
        UseCase->>Session: Create(ctx, scope, sessionInfo)
        Session-->>UseCase: Session created
        
        UseCase->>Role: Detail(ctx, scope, user.RoleID)
        Role-->>UseCase: Role details
    end
    
    %% Process Phase
    rect rgb(255, 255, 240)
        UseCase-->>API: Success
        API-->>Client: 200 OK
        Note over Client,API: Response: {user, role, token: {accessToken, refreshToken, expiresAt, sessionID}}
    end
```

## 7. Refresh Token Flow

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant UseCase
    participant Session
    participant UserRepo
    participant Role

    Client->>API: POST /api/v1/auth/refresh-token
    Note over Client,API: Headers: {Authorization: Bearer refreshToken}
    
    API->>UseCase: RefreshToken(ctx, scope, input)
    
    %% Verify Phase
    rect rgb(255, 240, 240)
        UseCase->>Session: GetByRefreshToken(ctx, scope, token)
        alt Session not found
            Session-->>UseCase: Error (Invalid token)
            UseCase-->>API: Error (Invalid token)
            API-->>Client: 401 Unauthorized
        end
        
        UseCase->>UseCase: Validate token expiry
        alt Token expired
            UseCase-->>API: Error (Token expired)
            API-->>Client: 401 Unauthorized
        end
    end
    
    %% Authenticate Phase
    rect rgb(240, 255, 240)
        UseCase->>UserRepo: GetOne(ctx, scope, session.UserID)
        UserRepo-->>UseCase: User details
        
        UseCase->>UseCase: Generate new tokens
        UseCase->>Session: Update(ctx, scope, {newTokens})
    end
    
    %% Fetch Phase
    rect rgb(240, 240, 255)
        UseCase->>Role: Detail(ctx, scope, user.RoleID)
        Role-->>UseCase: Role details
    end
    
    %% Process Phase
    rect rgb(255, 255, 240)
        UseCase-->>API: Success
        API-->>Client: 200 OK
        Note over Client,API: Response: {user, role, token: {accessToken, refreshToken, expiresAt, sessionID}}
    end
```

## 8. Logout Flow

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant UseCase
    participant Session

    Client->>API: POST /api/v1/auth/logout
    Note over Client,API: Headers: {Authorization: Bearer accessToken}
    
    API->>UseCase: Logout(ctx, scope, sessionID)
    
    %% Verify Phase
    rect rgb(255, 240, 240)
        UseCase->>Session: GetByID(ctx, scope, sessionID)
        alt Session not found
            Session-->>UseCase: Error (Session not found)
            UseCase-->>API: Error (Session not found)
            API-->>Client: 404 Not Found
        end
    end
    
    %% Process Phase
    rect rgb(255, 255, 240)
        UseCase->>Session: Delete(ctx, scope, sessionID)
        UseCase-->>API: Success
        API-->>Client: 200 OK
    end
```

## 9. Get User Profile Flow

```mermaid
sequenceDiagram
    participant Client
    participant API
    participant UseCase
    participant UserRepo
    participant Role

    Client->>API: GET /api/v1/auth/me
    Note over Client,API: Headers: {Authorization: Bearer accessToken}
    
    API->>UseCase: DetailMe(ctx, scope)
    
    %% Verify Phase
    rect rgb(255, 240, 240)
        Note over UseCase: User already authenticated via middleware
    end
    
    %% Fetch Phase
    rect rgb(240, 240, 255)
        UseCase->>UserRepo: GetOne(ctx, scope, userID)
        UserRepo-->>UseCase: User details
        
        UseCase->>Role: Detail(ctx, scope, user.RoleID)
        Role-->>UseCase: Role details
    end
    
    %% Process Phase
    rect rgb(255, 255, 240)
        UseCase-->>API: Success
        API-->>Client: 200 OK
        Note over Client,API: Response: {user, role}
    end
```

## Key Points

### Register API
- Check if email already exists
- Encrypt password before storing
- Create new user with unverified status
- No OTP generation during registration

### Send OTP API
- Verify user existence
- Decrypt and authenticate password
- Check OTP validity and generate new if needed
- Create email template with OTP details
- Send email via RabbitMQ

### Verify OTP API
- Verify user existence
- Verify OTP matches stored OTP
- Check OTP expiration
- Update user verification status

### Login API
- Verify user existence and verification status
- Decrypt and authenticate password
- Generate access token with JWT
- Generate refresh token using UUID
- Create session with token details
- Get user role information
- Support remember me functionality
- Return user, role and token information

### Social Login API
- Validate social provider (Facebook/Google/GitLab)
- Get OAuth2 configuration for the provider
- Generate authentication URL with:
  - Random state parameter
  - Offline access type
  - Account selection prompt
- Return authentication URL to client

### Social Callback API
- Validate social provider
- Get user info from OAuth2 provider
- Check if user exists
- Create new user if not exists
- Verify provider and providerID
- Generate token and session
- Get user role information
- Return user, role and token information

### Common Features
- Use RabbitMQ for asynchronous email delivery
- Password encryption/decryption
- Error handling and logging
- Use scope for access control management
- Email template generation

### VAFP Pattern
Each API flow follows the VAFP (Verify, Authenticate, Fetch, Process) pattern:

1. **Verify Phase** (Red)
   - Input validation
   - Existence checks
   - Error conditions
   - Data validation

2. **Authenticate Phase** (Green)
   - Security checks
   - Password/OTP verification
   - Token generation
   - Provider validation

3. **Fetch Phase** (Blue)
   - Data retrieval
   - Session creation
   - Role fetching
   - User information gathering

4. **Process Phase** (Yellow)
   - Response preparation
   - Final processing
   - Success/error handling
   - Response sending
