import { Component } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { DialogService } from 'src/app/features/select-business-dialog/dialog.service';
import { UserService } from 'src/app/services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  loginForm: FormGroup;
  errorMessage: string = '';

  constructor(
    private fb: FormBuilder,
    private router: Router,
    private dialogService: DialogService,
    private userService: UserService
  ) {
    this.loginForm = this.fb.group({
      email: ['', [Validators.required, Validators.email]],
      password: ['', Validators.required]
    });
  }

  onSubmit() {
    if (this.loginForm.valid) {
      const { email, password } = this.loginForm.value;

      this.userService.login(email, password).subscribe({
        next: (response) => {
          // Giả sử API trả về thông tin người dùng, bạn có thể lưu vào localStorage
          localStorage.setItem('isLoggedIn', 'true');
          localStorage.setItem('currentUser', JSON.stringify(response)); // Lưu response từ API nếu có
          
          // Mở dialog chọn doanh nghiệp
          this.dialogService.openBusinessSelectionDialog();
        },
        error: (error) => {
          this.errorMessage = error.error?.message || 'Email hoặc mật khẩu không đúng';
        }
      });
    }
  }
}