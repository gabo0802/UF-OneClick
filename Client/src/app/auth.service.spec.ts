import { HttpClientModule } from '@angular/common/http';
import { TestBed } from '@angular/core/testing';
import { RouterTestingModule } from '@angular/router/testing';

import { AuthService } from './auth.service';

describe('AuthService', () => {
  let service: AuthService;

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [],
      providers: [
        HttpClientModule, 
        RouterTestingModule,        
      ]
    });
    service = TestBed.inject(AuthService);    
  });

  it('Auth Service should be created', () => {
    expect(service).toBeTruthy();
  });


  it('visitor should not be logged in', () => {    

    expect(service.isLoggedIn()).not.toBeTruthy();
  });

  it('visitor logging in should be logged in', () => {
    
    service.userLogIn();
    expect(service.isLoggedIn()).toBeTruthy();
  })

  it('user should be logged in', () => {

    service.userLogIn();
    expect(service.isLoggedIn()).toBeTruthy();
  })

  it('user logging out should be logged out', () => {

    service.userLogIn();
    service.userLogOut();
    expect(service.isLoggedIn()).not.toBeTruthy();
  })

  it('returning logged in user with cookie should be logged in', () => {

    document.cookie = "currentUserID=";
    
    expect(service.isLoggedIn()).toBeTruthy();
    service.userLogOut();
  })

  
});
