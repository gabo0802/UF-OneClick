import { HttpClient } from "@angular/common/http"
import { LoginComponent } from "./login.component"
import { MatDialog } from '@angular/material/dialog';


describe('LoginComponent', () => {
    it('mounts Login', () => {
      cy.mount(LoginComponent, {
        providers: [{ provide: HttpClient, useValue: null }],
      })
    })    
  })