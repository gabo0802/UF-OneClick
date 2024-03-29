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
import { AdminGuard, AuthGuard, LogInGuard } from './auth-guard.guard';
import { AuthService } from './auth.service';
import { ApiService } from './api.service';
import { UsersComponent } from './dashboard/users/users.component';
import { AdminComponent } from './dashboard/admin/admin.component';
import { ProfileComponent } from './profile/profile.component';
import { FooterComponent } from './footer/footer.component';
import { SuccessComponent } from './dialogs/success/success.component';
import { DialogsService } from './dialogs.service';
import { ErrorComponent } from './dialogs/error/error.component';
import { UsernameFieldComponent } from './profile/input-fields/username-field/username-field.component';
import { EmailFieldComponent } from './profile/input-fields/email-field/email-field.component';
import { PasswordFieldComponent } from './profile/input-fields/password-field/password-field.component';
import { PasswordResetComponent } from './dialogs/password-reset/password-reset.component';
import { TimezoneFieldComponent } from './profile/input-fields/timezone-field/timezone-field.component';
import { DeleteAccountComponent } from './dialogs/delete-account/delete-account.component';
import { SubscriptionListComponent } from './dashboard/users/subscription-list/subscription-list.component';
import { AddSubscriptionComponent } from './dialogs/add-subscription/add-subscription.component';
import { AddInactiveSubscriptionComponent } from './dialogs/add-inactive-subscription/add-inactive-subscription.component';
import { ReportComponent } from './dashboard/users/report/report.component';
import { GraphComponent } from './dashboard/users/graph/graph.component';
import { NgChartsModule } from 'ng2-charts';
import { ResetComponent } from './dashboard/admin/reset/reset.component';
import { EntryComponent } from './dashboard/admin/entry/entry.component';
import { NewsletterComponent } from './dashboard/admin/newsletter/newsletter.component';

@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    LandingPageComponent,
    LoginComponent,
    SignupComponent,
    DashboardComponent,
    UsersComponent,
    AdminComponent,
    ProfileComponent,
    FooterComponent,
    SuccessComponent,
    ErrorComponent,
    UsernameFieldComponent,
    EmailFieldComponent,
    PasswordFieldComponent,
    PasswordResetComponent,
    TimezoneFieldComponent,
    DeleteAccountComponent,
    SubscriptionListComponent,
    AddSubscriptionComponent,
    AddInactiveSubscriptionComponent,
    ReportComponent,
    GraphComponent,
    ResetComponent,
    EntryComponent,
    NewsletterComponent,
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FormsModule,
    ReactiveFormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
   
    MaterialDesignModule,
        NgChartsModule
  ],
  providers: [AuthService, AuthGuard, ApiService, LogInGuard, DialogsService, AdminGuard],
  bootstrap: [AppComponent]
})
export class AppModule { }
