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

    it('Add Subscription Button has text Add Subscription', () => {
        cy.mount(SubscriptionListComponent, {
          imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],         
          providers: [ApiService, AuthGuard, AuthService, Router]
        })

        cy.get('[class="buttons-container"]').get('button').should('have.text', 'Add Subscription')
    })

})