import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { ScheduleService } from 'src/app/services/schedule.service';
import { DatePipe } from '@angular/common';
import { MatDialog } from '@angular/material/dialog';
import { FaceRegistrationDialogComponent } from '../../face-resgiter/face-register.component';

interface Schedule {
  date: string;
  dayName: string;
  shift: string;
  status: string;
}

@Component({
  selector: 'app-dashboard',
  templateUrl: './dashboard.component.html',
  styleUrls: ['./dashboard.component.css'],
  providers: [DatePipe]
})
export class DashboardComponent implements OnInit {
  schedule: Schedule[] = [];
  currentMonth: string;
  businessId: string = '';
  businessName = 'Business 1';
  position = 'Quản lý';
  isLoading: boolean = false;

  constructor(
    private route: ActivatedRoute,
    private scheduleService: ScheduleService,
    private datePipe: DatePipe,
    private dialog: MatDialog,
  ) {
    this.currentMonth = this.datePipe.transform(new Date(), 'MM-yyyy') || '12-2025';
  }

  ngOnInit() {
    this.businessId = this.route.snapshot.paramMap.get('businessId') || '';
    this.loadSchedule();
  }

  loadSchedule() {
    this.isLoading = true;
    this.scheduleService.getSchedule(this.businessId, this.currentMonth)
      .subscribe({
        next: (data) => {
          this.schedule = data;
          this.isLoading = false;
        },
        error: (err) => {
          console.error(err);
          this.isLoading = false;
        }
      });
  }

  changeMonth(offset: number): void {
    // Tạo date object từ currentMonth
    const [month, year] = this.currentMonth.split('-').map(Number);
    const date = new Date(year, month - 1 + offset, 1);
    
    // Cập nhật currentMonth mới
    this.currentMonth = this.datePipe.transform(date, 'MM-yyyy') || '12-2025';
    
    // Load lại dữ liệu cho tháng mới
    this.loadSchedule();
  }
    // Mở dialog đăng ký khuôn mặt
  openFaceRegistration(): void {
    const dialogRef = this.dialog.open(FaceRegistrationDialogComponent, {
      width: '500px',
      data: { userId: '123' } // Truyền dữ liệu nếu cần
    });

    dialogRef.afterClosed().subscribe(result => {
      if (result) {
        console.log('Đã đăng ký khuôn mặt:', result);
        // Có thể thêm thông báo thành công
      }
    });
  }
}