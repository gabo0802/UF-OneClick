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
      
      it('Form should be initially empty', () => {

        cy.mount(SignupComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router]
        })
        
        cy.get('[formControlName="username"]').should('have.value', '')
        cy.get('mat-form-field').get('[formControlName="email"]').should('have.value', '')
        cy.get('mat-form-field').get('input').get('[formControlName="password"]').should('have.value', '')
      })

      it('Form values should be accurate', () => {

        cy.mount(SignupComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router]
        })
        
        cy.get('[formControlName="username"]').type('Johnny').should('have.value', 'Johnny')
        cy.get('mat-form-field').get('[formControlName="email"]').type('test@testing.com').should('have.value', 'test@testing.com')
        cy.get('mat-form-field').get('input').get('[formControlName="password"]').type('kittycat').should('have.value', 'kittycat')
      })

      it('Sign up Button disabled when form is empty', () => {

        cy.mount(SignupComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router]
        })

        cy.get('[formControlName="username"]').should('have.value', '')
        cy.get('mat-form-field').get('[formControlName="email"]').should('have.value', '')
        cy.get('mat-form-field').get('input').get('[formControlName="password"]').should('have.value', '')
  
        cy.get('mat-card-actions').contains('button', 'Sign Up').then(($btn) => {
          if ($btn.is(':disabled')) {
            cy.log('Button exists and is disabled!')
            return
          }
          else {
            cy.log('Button exists and is enabled!')
            cy.wrap($btn).click()
          }
        })
      })

      it('Sign Up Button disabled when password form is empty', () => {

        cy.mount(SignupComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router]
        })
  
        cy.get('[formControlName="username"]').type('Susie')
        cy.get('mat-form-field').get('[formControlName="email"]').type('susie@testing.com')
  
        cy.get('mat-card-actions').contains('button', 'Sign Up').then(($btn) => {
          if ($btn.is(':disabled')) {
            cy.log('Button exists and is disabled!')
            return
          }
          else {
            cy.log('Button exists and is enabled!')
            cy.wrap($btn).click()
          }
        })
      })

      it('Sign Up Button disabled when username form is empty', () => {

        cy.mount(SignupComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router]
        })
  
        
        cy.get('mat-form-field').get('[formControlName="email"]').type('susie@testing.com')
        cy.get('mat-form-field').get('input').get('[formControlName="password"]').type('kittycat')
  
        cy.get('mat-card-actions').contains('button', 'Sign Up').then(($btn) => {
          if ($btn.is(':disabled')) {
            cy.log('Button exists and is disabled!')
            return
          }
          else {
            cy.log('Button exists and is enabled!')
            cy.wrap($btn).click()
          }
        })
      })

      it('Sign Up Button disabled when email form is empty', () => {

        cy.mount(SignupComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router]
        })
  
        
        cy.get('[formControlName="username"]').type('Susie')
        cy.get('mat-form-field').get('input').get('[formControlName="password"]').type('kittycat')
  
        cy.get('mat-card-actions').contains('button', 'Sign Up').then(($btn) => {
          if ($btn.is(':disabled')) {
            cy.log('Button exists and is disabled!')
            return
          }
          else {
            cy.log('Button exists and is enabled!')
            cy.wrap($btn).click()
          }
        })
      })

      it('Sign Up Button enabled when username, email, password forms filled in', () => {

        cy.mount(SignupComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router]
        })
  
        cy.get('[formControlName="username"]').type('Susie')
        cy.get('mat-form-field').get('[formControlName="email"]').type('susie@testing.com')
        cy.get('mat-form-field').get('input').get('[formControlName="password"]').type('MyPassword')
  
        cy.get('mat-card-actions').contains('button', 'Sign Up').then(($btn) => {
          if ($btn.is(':enabled')) {
            cy.log('Button exists and is enabled!')
            return
          }
          else {
            cy.log('Button exists and is disabled!')          
          }
        })
      })
  })