import { Component, ViewChild } from '@angular/core';
import { FormBuilder, FormGroup, Validators, NgForm, FormControl } from '@angular/forms';
import { Router } from '@angular/router';
import { UserService } from '../../../../services/user.service';
import { RegisterParam } from '../../../../params/register.param';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  @ViewChild('registerForm') registerFormDirective!: NgForm;
  registerForm: FormGroup;
  otpControl: FormControl;
  isLoading = false;
  errorMessage = '';
  showOtpInput = false;

  constructor(
    private fb: FormBuilder, 
    private router: Router,
    private userService: UserService
  ) {
    this.registerForm = this.fb.group({
      fullName: ['', [Validators.required, Validators.minLength(3)]],
      email: ['', [Validators.required, Validators.email]],
      password: ['', [Validators.required, Validators.minLength(6)]],
      confirmPassword: ['', Validators.required],
      agreeTerms: [false, Validators.requiredTrue]
    }, { 
      validator: this.passwordMatchValidator,
      updateOn: 'blur'
    });

    this.otpControl = this.fb.control('', [Validators.required, Validators.minLength(6)]);
  }

  passwordMatchValidator(form: FormGroup) {
    const password = form.get('password');
    const confirmPassword = form.get('confirmPassword');
    
    if (password?.value !== confirmPassword?.value) {
      confirmPassword?.setErrors({ mismatch: true });
    } else {
      confirmPassword?.setErrors(null);
    }
    return null;
  }

  onSubmit() {
    if (this.registerForm.valid) {
      this.isLoading = true;
      this.errorMessage = '';

      const registerParam: RegisterParam = {
        username: this.registerForm.value.fullName,
        email: this.registerForm.value.email,
        password: this.registerForm.value.password
      };

      // Gọi API đăng ký
      this.userService.register(registerParam).subscribe({
        next: () => {
          this.showOtpInput = true; // Chuyển sang bước nhập OTP
          this.isLoading = false;
        },
        error: (error) => {
          this.errorMessage = error.error?.message || 'Đăng ký thất bại. Vui lòng thử lại sau.';
          this.isLoading = false;
        }
      });
    } else {
      Object.keys(this.registerForm.controls).forEach(key => {
        const control = this.registerForm.get(key);
        control?.markAsTouched();
      });
    }
  }

  onSendOtp() {
    this.isLoading = true;
    this.errorMessage = '';

    // Gọi API gửi OTP
    this.userService.sendOtp(this.registerForm.value.email, this.registerForm.value.password).subscribe({
      next: () => {
        this.isLoading = false;
        this.errorMessage = 'OTP đã được gửi đến email của bạn.';
      },
      error: (error) => {
        this.errorMessage = error.error?.message || 'Gửi OTP thất bại. Vui lòng thử lại sau.';
        this.isLoading = false;
      }
    });
  }

  onVerifyOtp() {
    if (this.otpControl.invalid) {
      this.errorMessage = 'Vui lòng nhập mã OTP hợp lệ (tối thiểu 6 ký tự).';
      return;
    }

    this.isLoading = true;
    this.errorMessage = '';

    // Gọi API xác minh OTP
    this.userService.verifyOtp(this.registerForm.value.email, this.otpControl.value).subscribe({
      next: () => {
        this.router.navigate(['/login']);
        this.isLoading = false;
      },
      error: (error) => {
        this.errorMessage = error.error?.message || 'Xác minh OTP thất bại. Vui lòng thử lại.';
        this.isLoading = false;
      }
    });
  }
}