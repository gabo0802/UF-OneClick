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
import { ScrollBarComponent } from './scroll-bar/scroll-bar.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { SubscriptionListComponent } from './dashboard/subscription-list/subscription-list.component';
import { SubscriptionDetailComponent } from './dashboard/subscription-list/subscription-detail/subscription-detail.component';
import { SignupMessageComponent } from './signup/signup-message/signup-message.component';



@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    LandingPageComponent,
    LoginComponent,
    SignupComponent,
    ScrollBarComponent,
    DashboardComponent,
    SubscriptionListComponent,
    SubscriptionDetailComponent,
    SignupMessageComponent
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
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
