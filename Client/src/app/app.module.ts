import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialDesignModule } from './material-design/material-design.module';
import { HeaderComponent } from './header/header.component';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { LoginComponent } from './login/login.component';
import { SignupComponent } from './signup/signup.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { SignupMessageComponent } from './signup/signup-message/signup-message.component';
import { LoginMessageComponent } from './login/login-message/login-message.component';
import { AuthGuard, LogInGuard } from './auth-guard.guard';
import { AuthService } from './auth.service';
import { ApiService } from './api.service';
import { UsersComponent } from './dashboard/users/users.component';
import { AdminComponent } from './dashboard/admin/admin.component';
import { ProfileComponent } from './profile/profile.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    LandingPageComponent,
    LoginComponent,
    SignupComponent,
    DashboardComponent,
    SignupMessageComponent,
    LoginMessageComponent,
    UsersComponent,
    AdminComponent,
    ProfileComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
   
    MaterialDesignModule
  ],
  providers: [AuthService, AuthGuard, ApiService, LogInGuard],
  bootstrap: [AppComponent]
})
export class AppModule { }
