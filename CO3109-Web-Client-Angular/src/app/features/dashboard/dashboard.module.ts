import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { DashboardComponent } from './component/dashboard.component';
import { DashboardRoutingModule } from './dashboard-routing.module';
import { MatDialogModule } from '@angular/material/dialog';
import { FaceRegistrationDialogComponent } from '../face-resgiter/face-register.component';




@NgModule({
  declarations: [
    DashboardComponent,
    FaceRegistrationDialogComponent // Thêm dialog component
  ],
  imports: [
    CommonModule,
    DashboardRoutingModule,
    MatDialogModule, // Thêm module dialog
  ]
})
export class DashboardModule { }