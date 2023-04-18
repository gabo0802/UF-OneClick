import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from "@angular/platform-browser"
import { BrowserAnimationsModule } from "@angular/platform-browser/animations"
import { Router } from "@angular/router"
import { ApiService } from "src/app/api.service"
import { AppRoutingModule } from "src/app/app-routing.module"
import { AuthGuard } from "src/app/auth-guard.guard"
import { AuthService } from "src/app/auth.service"
import { MaterialDesignModule } from "src/app/material-design/material-design.module"
import { SubscriptionListComponent } from "./subscription-list.component"



describe('SubscriptionListComponent', () => {

    it('mounts', () => {
      cy.mount(SubscriptionListComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],         
        providers: [ApiService, AuthGuard, AuthService, Router]
      })
    })

    it('Add Active Subscription Button has text Add Active Subscription', () => {
        cy.mount(SubscriptionListComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],         
          providers: [ApiService, AuthGuard, AuthService, Router]
        })

        cy.get('[class="buttons-container mat-elevation-z2"]').get('button').first().should('have.text', 'Add Active Subscription')
    })

    it('Add Inactive Subscription Button has text Add Inactive Subscription', () => {
      cy.mount(SubscriptionListComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],         
        providers: [ApiService, AuthGuard, AuthService, Router]
      })

      cy.get('[class="buttons-container mat-elevation-z2"]').get('button').next().should('have.text', 'Add Inactive Subscription')
  })

})