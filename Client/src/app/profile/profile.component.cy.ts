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
import { DialogsService } from "../dialogs.service"
import { MaterialDesignModule } from "../material-design/material-design.module"
import { ProfileComponent } from "./profile.component"
import { ErrorComponent } from "../dialogs/error/error.component"
import { DeleteAccountComponent } from "../dialogs/delete-account/delete-account.component"

describe('ProfileComponent', () => {    

    it('mounts Profile', () => {
      cy.mount(ProfileComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
        declarations: [ErrorComponent, DeleteAccountComponent],
        providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
      })
    })

    it('Correct heading for Profile Component', () => {

      cy.mount(ProfileComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
        declarations: [ErrorComponent, DeleteAccountComponent],
        providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
      })

      cy.get('mat-card-header').get('h3').should('have.text', 'Profile Information')

    })

    it('Correct delete profile button text', () => {

      cy.mount(ProfileComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
        declarations: [ErrorComponent, DeleteAccountComponent],
        providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
      })

      cy.get('mat-card-actions').get('button').first().should('have.text', 'Delete Profile')
    })

    it('Correct back button text', () => {

      cy.mount(ProfileComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
        declarations: [ErrorComponent, DeleteAccountComponent],
        providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
      })

      cy.get('mat-card-actions').get('button').first().next().should('have.text', 'Back')
    })

    it('Delete profile button Clickable', () => {

      cy.mount(ProfileComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
        declarations: [ErrorComponent, DeleteAccountComponent],
        providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
      })

      cy.get('mat-card-actions').get('button').first().click()
    })
})