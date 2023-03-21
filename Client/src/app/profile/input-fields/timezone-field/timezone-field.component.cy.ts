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
import { TimezoneFieldComponent } from "./timezone-field.component"

describe('Time Zone-Field Component', () => {    

    it('mounts Time Zone-Field Component', () => {
      cy.mount(TimezoneFieldComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
        declarations: [],
        providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
      })
    });

    it('Label contains Time Zone:', () => {
        cy.mount(TimezoneFieldComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
        })

        cy.get('[class="label"]').should('have.text', 'Time Zone:')
    });

    it('Edit button contains text Edit', () => {
        cy.mount(TimezoneFieldComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
        })

        cy.get('button').first().should('have.text', 'Edit')
    });

    it('Edit button should change text to Save when clicked', () => {
        cy.mount(TimezoneFieldComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule, ReactiveFormsModule],
          declarations: [],
          providers: [ApiService, AuthService, AuthGuard, MatDialog, Router, DialogsService]
        })

        cy.get('button').first().click()
        cy.get('button').first().should('have.text', 'Save')        
    });
});