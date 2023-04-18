import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { SignupComponent } from './signup/signup.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { AdminGuard, AuthGuard, LogInGuard } from './auth-guard.guard';
import { UsersComponent } from './dashboard/users/users.component';
import { ProfileComponent } from './profile/profile.component';
import { AdminComponent } from './dashboard/admin/admin.component';
import { ResetComponent } from './dashboard/admin/reset/reset.component';
import { EntryComponent } from './dashboard/admin/entry/entry.component';
import { NewsletterComponent } from './dashboard/admin/newsletter/newsletter.component';

const routes: Routes = [
  {path: '', canActivate: [LogInGuard],  component: LandingPageComponent, pathMatch: 'full'},
  {path: 'login', canActivate: [LogInGuard], component: LoginComponent},
  {path: 'signup', canActivate: [LogInGuard], component: SignupComponent},
  {path: 'users', canActivate: [AuthGuard], component: DashboardComponent, children: [
    {path: '', component: UsersComponent},
    {path: 'profile', component: ProfileComponent},
  ]},
  {path: 'admin', canActivate: [AdminGuard], component: AdminComponent, children: [
    {path: '', component: EntryComponent},
    {path: 'reset', component: ResetComponent},
    {path: 'newsletter', component: NewsletterComponent}
  ]}  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
