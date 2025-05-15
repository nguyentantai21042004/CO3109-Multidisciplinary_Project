import { Injectable } from '@angular/core';
import { Observable, of } from 'rxjs';
import { delay } from 'rxjs/operators';

interface Schedule {
  date: string;
  dayName: string;
  shift: string;
  status: string;
}

@Injectable({ providedIn: 'root' })
export class ScheduleService {
  constructor() {}

  getSchedule(businessId: string, month: string): Observable<Schedule[]> {
    // Parse month (format: MM-yyyy)
    const [monthNumber, year] = month.split('-').map(Number);
    const daysInMonth = new Date(year, monthNumber, 0).getDate();
    
    // Tạo dữ liệu giả cho tháng hiện tại
    const mockSchedule: Schedule[] = [];
    
    for (let day = 1; day <= daysInMonth; day++) {
      const date = new Date(year, monthNumber - 1, day);
      const dayName = this.getDayName(date.getDay());
      
      // Random shift và status (giả lập dữ liệu)
      const shift = this.generateRandomShift();
      const status = this.generateRandomStatus(date);
      
      mockSchedule.push({
        date: `${day.toString().padStart(2, '0')}/${monthNumber.toString().padStart(2, '0')}`,
        dayName: dayName,
        shift: shift,
        status: status
      });
    }

    // Trả về dữ liệu giả với delay để giả lập API call
    return of(mockSchedule).pipe(delay(500));
  }

  private getDayName(dayIndex: number): string {
    const days = ['Chủ Nhật', 'Thứ Hai', 'Thứ Ba', 'Thứ Tư', 'Thứ Năm', 'Thứ Sáu', 'Thứ Bảy'];
    return days[dayIndex];
  }

  private generateRandomShift(): string {
    const shifts = ['Ca sáng (8h-12h)', 'Ca chiều (13h-17h)', 'Cả ngày (8h-17h)', 'Nghỉ'];
    return shifts[Math.floor(Math.random() * shifts.length)];
  }

  private generateRandomStatus(date: Date): string {
    const day = date.getDay();
    const isWeekend = day === 0 || day === 6; // Chủ Nhật hoặc Thứ 7
    
    if (isWeekend) {
      return 'Nghỉ';
    } else {
      const statuses = ['Đã chấm công', 'Chưa chấm công', 'Đi trễ', 'Vắng'];
      return statuses[Math.floor(Math.random() * statuses.length)];
    }
  }
}