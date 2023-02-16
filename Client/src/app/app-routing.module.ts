import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { SignupComponent } from './signup/signup.component';
import { DashboardComponent } from './dashboard/dashboard.component';
import { AuthGuard, LogInGuard } from './auth-guard.guard';

const routes: Routes = [
  {path: '', canActivate: [LogInGuard],  component: LandingPageComponent, pathMatch: 'full'},
  {path: '', canActivate: [AuthGuard], component: DashboardComponent, pathMatch: 'full'},
  {path: 'login', canActivate: [LogInGuard], component: LoginComponent},
  {path: 'signup', canActivate: [LogInGuard], component: SignupComponent},
  {path: 'users', canActivate: [AuthGuard], component: DashboardComponent}  
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
