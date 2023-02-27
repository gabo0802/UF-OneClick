import { HttpClientModule } from "@angular/common/http"
import { ReactiveFormsModule } from "@angular/forms"
import { MatDialog } from "@angular/material/dialog"
import { BrowserModule } from "@angular/platform-browser"
import { BrowserAnimationsModule } from "@angular/platform-browser/animations"
import { Router } from "@angular/router"
import { ApiService } from "../api.service"
import { AppRoutingModule } from "../app-routing.module"
import { AuthGuard } from "../auth-guard.guard"
import { AuthService } from "../auth.service"
import { MaterialDesignModule } from "../material-design/material-design.module"
import { SignupComponent } from "./signup.component"



describe('SignUpComponent', () => {
    it('mounts Sign Up', () => {
      cy.mount(SignupComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
        declarations: [SignupComponent],
        providers: [ApiService, AuthService, AuthGuard, MatDialog, Router]
      })
    })
    
    it('Able to fill In Sign Up form', () => {
        cy.mount(SignupComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [SignupComponent],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router]
        })
        
        cy.get('mat-form-field').get('[formControlName="username"]').type('Username')
        cy.get('mat-form-field').get('[formControlName="email"]').type('test@123.com')
        cy.get('mat-form-field').get('input').get('[formControlName="password"]').type('password')       
        
      }) 
  })