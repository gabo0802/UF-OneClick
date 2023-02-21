import { HeaderComponent } from "./header.component";

describe('HeaderComponent', () => {
    it('mounts', () => {
      cy.mount(HeaderComponent)
    })

    it('Login and Sign Up button text', () => {
        cy.mount(HeaderComponent)
        cy.get('[routerLink="/login"]').should('have.text', 'Login')
        cy.get('[routerLink="/signup"]').should('have.text', 'Sign Up')
    })
  })