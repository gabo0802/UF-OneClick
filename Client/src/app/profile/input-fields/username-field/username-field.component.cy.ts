import { HttpClientModule } from "@angular/common/http"
import { ReactiveFormsModule } from "@angular/forms"
import { MatDialog } from "@angular/material/dialog"
import { BrowserModule } from "@angular/platform-browser"
import { BrowserAnimationsModule } from "@angular/platform-browser/animations"
import { Router } from "@angular/router"
import { ApiService } from "src/app/api.service"
import { AppRoutingModule } from "src/app/app-routing.module"
import { AuthGuard } from "src/app/auth-guard.guard"
import { AuthService } from "src/app/auth.service"
import { DialogsService } from "src/app/dialogs.service"
import { MaterialDesignModule } from "src/app/material-design/material-design.module"
import { UsernameFieldComponent } from "./username-field.component"

describe('Username-Field Component', () => {    

    it('mounts Username-Field Component', () => {
      cy.mount(UsernameFieldComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
        declarations: [],
        providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
      })
    });

    it('Label contains Username', () => {
        cy.mount(UsernameFieldComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
        })

        cy.get('label').should('have.text', 'Username:')
    });

    it('Edit button contains text Edit', () => {
        cy.mount(UsernameFieldComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
        })

        cy.get('button').first().should('have.text', 'Edit')
    });

    it('Edit button should change text to Save when clicked', () => {
        cy.mount(UsernameFieldComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
        })

        cy.get('button').first().click()
        cy.get('button').first().should('have.text', 'Save')        
    });

    it('Valid Input into form enables button', () => {

        cy.mount(UsernameFieldComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
        })

        cy.get('button').first().click()
        cy.get('input').type('OrigamiBear')
        cy.get('button').first().then(($btn) => {

            if ($btn.is(':enabled')) {
                cy.log('Valid Input and is enabled!')
                return
              }
              else {
                cy.log('Invalid Input and is disabled!')          
              }
        })       
    });
});