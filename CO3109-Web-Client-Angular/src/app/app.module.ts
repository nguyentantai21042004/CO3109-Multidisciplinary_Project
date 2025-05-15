import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatDialogModule } from '@angular/material/dialog';

import { AppRoutingModule } from './app-routing.module';
import { HomeComponent } from './features/home/components/home.component';
import { AppComponent } from './app.component';
import { LoginModule } from './features/auth/login/login.module';
import { SelectBusinessDialogComponent } from './features/select-business-dialog/component/select-business-dialog.component';
import { DialogService } from './features/select-business-dialog/dialog.service';
import { RegisterModule } from './features/auth/register/register.module';
import { HttpClientModule } from '@angular/common/http';
import { NotFound404Component } from './features/not-found404/not-found404.component';

@NgModule({
  declarations: [
    AppComponent,
    SelectBusinessDialogComponent,
    NotFound404Component,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    LoginModule,
    MatDialogModule,
    RegisterModule,
    HttpClientModule,
  ],
  providers: [DialogService],
  bootstrap: [
    AppComponent,
  ]
})
export class AppModule { }
