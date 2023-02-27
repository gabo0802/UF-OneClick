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
import { LoginMessageComponent } from "./login-message/login-message.component"
import { LoginComponent } from "./login.component"


describe('LoginComponent', () => {
    it('mounts Login', () => {
      cy.mount(LoginComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
        declarations: [LoginMessageComponent],
        providers: [ApiService, AuthService, AuthGuard, MatDialog, Router]
      })
    })
    
    it('Able to fill In Login form', () => {
        cy.mount(LoginComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [LoginMessageComponent],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router]
        })
        
        cy.get('[formControlName="username"]').type("Username")
        cy.get('mat-form-field').get('input').get('[formControlName="password"]').type("password")
        
      }) 
  })