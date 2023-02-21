import { LandingPageComponent } from "./landing-page.component" 

describe('LandingPageComponent', () => {
    it('mounts', () => {
      cy.mount(LandingPageComponent)
    })

    it('Testing card title', () => {
        cy.mount(LandingPageComponent)
        cy.get('mat-card-title').should('have.text', 'What is OneClick?')
    })
  })