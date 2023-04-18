import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from "@angular/platform-browser"
import { BrowserAnimationsModule } from "@angular/platform-browser/animations"
import { Router } from "@angular/router"
import { ApiService } from "src/app/api.service"
import { AppRoutingModule } from "src/app/app-routing.module"
import { AuthGuard } from "src/app/auth-guard.guard"
import { AuthService } from "src/app/auth.service"
import { MaterialDesignModule } from "src/app/material-design/material-design.module"
import { GraphComponent } from "./graph.component"



describe('GraphComponent', () => {

    it('mounts', () => {
      cy.mount(GraphComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],         
        providers: [ApiService, AuthGuard, AuthService, Router]
      })
    })

    it('has correct title', () => {
        cy.mount(GraphComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],         
          providers: [ApiService, AuthGuard, AuthService, Router]
        })

        cy.get('mat-card-title').first().should('have.text', 'Total Monthly Payments')        
    })

    it('Start Date has correct label of Start Date:', () => {
      cy.mount(GraphComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],         
        providers: [ApiService, AuthGuard, AuthService, Router]
      })

      cy.get('[class="form-container"]').get('label').first().should('have.text', 'Start Date:')
    })

    it('End Date has correct label of End Date:', () => {
      cy.mount(GraphComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],         
        providers: [ApiService, AuthGuard, AuthService, Router]
      })

      cy.get('[class="form-container"]').get('[class="labels"]').eq(1).should('have.text', 'End Date:')
    })

    it('Able to type in Start Year input', () => {
      cy.mount(GraphComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],         
        providers: [ApiService, AuthGuard, AuthService, Router]
      })

      cy.get('[class="form-container"]').get('mat-form-field').eq(1).type('2018')
    })

    it('Able to type in End Year input', () => {
      cy.mount(GraphComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],         
        providers: [ApiService, AuthGuard, AuthService, Router]
      })

      cy.get('[class="form-container"]').get('mat-form-field').eq(3).type('2021')
    })
})