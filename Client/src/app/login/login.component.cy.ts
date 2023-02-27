import { HttpClientModule } from "@angular/common/http"
import { BrowserModule } from "@angular/platform-browser"
import { BrowserAnimationsModule } from "@angular/platform-browser/animations"
import { ApiService } from "../api.service"
import { AppRoutingModule } from "../app-routing.module"
import { AuthGuard } from "../auth-guard.guard"
import { AuthService } from "../auth.service"
import { MaterialDesignModule } from "../material-design/material-design.module"
import { LoginMessageComponent } from "./login-message/login-message.component"
import { LoginComponent } from "./login.component"


describe('LoginComponent', () => {
    it('mounts Login', () => {
      cy.mount(LoginComponent, {
        imports: [HttpClientModule, MaterialDesignModule, BrowserAnimationsModule, BrowserModule, AppRoutingModule],
        declarations: [LoginMessageComponent],
        providers: [ApiService, AuthService, AuthGuard]
      })
    })    
  })