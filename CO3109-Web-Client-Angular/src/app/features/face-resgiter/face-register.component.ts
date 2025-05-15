import { Component, Inject } from '@angular/core';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

@Component({
  selector: 'app-face-registration-dialog',
  template: `
    <h2 mat-dialog-title>Đăng ký khuôn mặt</h2>
    <mat-dialog-content>
      <div class="camera-preview">
        <div class="camera-placeholder">
          <svg width="48" height="48" viewBox="0 0 24 24" fill="#666">
            <path d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm0 18c-4.41 0-8-3.59-8-8s3.59-8 8-8 8 3.59 8 8-3.59 8-8 8zm5-13h-2v2h2v2h-2v2h-2v-2h-2v2H9v-2H7v-2h2V9H7V7h2V5h2v2h2V5h2v2z"/>
          </svg>
          <p>Khu vực hiển thị camera</p>
        </div>
        <button mat-button (click)="captureFace()" color="primary">
          <svg width="16" height="16" viewBox="0 0 24 24" fill="currentColor" style="margin-right: 8px;">
            <path d="M4 4h3l2-2h6l2 2h3a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2m8 3a5 5 0 0 0-5 5a5 5 0 0 0 5 5a5 5 0 0 0 5-5a5 5 0 0 0-5-5m0 2a3 3 0 0 1 3 3a3 3 0 0 1-3 3a3 3 0 0 1-3-3a3 3 0 0 1 3-3z"/>
          </svg>
          Chụp ảnh
        </button>
      </div>
    </mat-dialog-content>
    <mat-dialog-actions align="end">
      <button mat-button (click)="onCancel()">Hủy</button>
      <button mat-button (click)="onSubmit()" color="primary" [disabled]="!isCaptured">
        Xác nhận
      </button>
    </mat-dialog-actions>
  `,
  styles: [`
    .camera-preview {
      text-align: center;
      padding: 1rem;
    }
    .camera-placeholder {
      border: 2px dashed #ccc;
      padding: 2rem;
      margin-bottom: 1rem;
      border-radius: 8px;
    }
  `]
})
export class FaceRegistrationDialogComponent {
  isCaptured = false;

  constructor(
    public dialogRef: MatDialogRef<FaceRegistrationDialogComponent>,
    @Inject(MAT_DIALOG_DATA) public data: any
  ) {}

  captureFace(): void {
    this.isCaptured = true;
    console.log('Đã chụp khuôn mặt (Mock)');
  }

  onSubmit(): void {
    this.dialogRef.close({ success: true });
  }

  onCancel(): void {
    this.dialogRef.close();
  }
}