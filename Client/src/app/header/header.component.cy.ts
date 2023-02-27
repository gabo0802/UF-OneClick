import { HttpClientModule } from "@angular/common/http";
import { BrowserModule } from "@angular/platform-browser";
import { BrowserAnimationsModule } from "@angular/platform-browser/animations";
import { ApiService } from "../api.service";
import { AppRoutingModule } from "../app-routing.module";
import { MaterialDesignModule } from "../material-design/material-design.module";
import { HeaderComponent } from "./header.component";

describe('HeaderComponent', () => {
    it('mounts', () => {
      cy.mount(HeaderComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],        
        providers: [ApiService]
      })
    })

    it('Login button text', () => {
        cy.mount(HeaderComponent)
        cy.get('[routerLink="/login"]').within(() => {
          cy.get('button').should('have.text', 'Login')})        
    })

    it('Click login button', ()=>{
      cy.mount(HeaderComponent)
      cy.get('[routerLink="/login"]').within( ()=> {
        cy.get('button').click()
      })
    })

    it('Signup button text', () => {
      cy.mount(HeaderComponent)
      cy.get('[routerLink="/signup"]').within(() => {
        cy.get('button').should('have.text', 'Sign Up')})        
  })

    it('Click Sign up button', ()=>{
      cy.mount(HeaderComponent)
      cy.get('[routerLink="/signup"]').within( ()=> {
        cy.get('button').click()
      })
    })
  })