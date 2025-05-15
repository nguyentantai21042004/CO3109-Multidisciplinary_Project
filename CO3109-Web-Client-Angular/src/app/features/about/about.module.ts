import { NgModule } from '@angular/core';
import { AboutRoutingModule } from './about-routing.module';
import { AboutComponent } from './components/about/about.component';
import { CommonModule } from '@angular/common';
import { TestComponent } from './components/test/test.component';

@NgModule({
  declarations: [
    AboutComponent,
    TestComponent
],
  imports: [
    CommonModule,
    AboutRoutingModule,
  ],
})
export class AboutModule { }    
