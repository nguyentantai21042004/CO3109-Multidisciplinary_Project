import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { DashboardComponent } from './component/dashboard.component';

const routes: Routes = [
  { 
    path: ':businessId', 
    component: DashboardComponent,
    data: { title: 'Dashboard' }
  }
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class DashboardRoutingModule { }
