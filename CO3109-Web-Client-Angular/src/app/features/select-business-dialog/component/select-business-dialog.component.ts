import { Component, EventEmitter, Output } from '@angular/core';
import { Router } from '@angular/router'; // Đã import đúng

interface Business {
  id: string;
  name: string;
  logo: string;
  position: string;
}

@Component({
  selector: 'app-select-business-dialog',
  templateUrl: './select-business-dialog.component.html',
  styleUrls: ['./select-business-dialog.component.css']
})
export class SelectBusinessDialogComponent {
  @Output() businessSelected = new EventEmitter<Business>();

  businesses: Business[] = [
    {
      id: '1',
      name: 'Doanh nghiệp 1',
      logo: 'assets/images/business1.jpg',
      position: 'Quản Lý'
    },
    {
      id: '2',
      name: 'Doanh nghiệp 2',
      logo: 'assets/images/business1.jpg',
      position: 'Nhân viên'
    },
    {
      id: '3',
      name: 'Doanh nghiệp 3',
      logo: 'assets/images/business1.jpg',
      position: 'Nhân viên'
    },
    {
      id: '4',
      name: 'Doanh nghiệp 4',
      logo: 'assets/images/business1.jpg',
      position: 'Nhân viên'
    },
    {
      id: '5',
      name: 'Doanh nghiệp 5',
      logo: 'assets/images/business1.jpg',
      position: 'Nhân viên'
    },
    {
      id: '6',
      name: 'Doanh nghiệp 6',
      logo: 'assets/images/business1.jpg',
      position: 'Nhân viên'
    },
    {
      id: '7',
      name: 'Doanh nghiệp 7',
      logo: 'assets/images/business1.jpg',
      position: 'Nhân viên'
    },
    {
      id: '8',
      name: 'Doanh nghiệp 8',
      logo: 'assets/images/business1.jpg',
      position: 'Nhân viên'
    },
    {
      id: '9',
      name: 'Doanh nghiệp 9',
      logo: 'assets/images/business1.jpg',
      position: 'Nhân viên'
    },
    {
      id: '10',
      name: 'Doanh nghiệp 10',
      logo: 'assets/images/business1.jpg',
      position: 'Nhân viên'
    }
  ];
  constructor(private router: Router) {}

  selectBusiness(business: Business) {
    this.businessSelected.emit(business);
    this.router.navigate(['/dashboard', business.id]);
    console.log('Navigating to:', business.id);
  }
  setDefaultImage(event: Event) {
    const img = event.target as HTMLImageElement;
    img.src = 'assets/images/default-business.png';
    img.onerror = null;
  }
}