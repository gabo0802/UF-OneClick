import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { PasswordResetComponent } from './password-reset/password-reset.component'
import { MatDialog } from '@angular/material/dialog';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {

  constructor(private router: Router, private dialog: MatDialog) {}

  hide: boolean = true;
  passwordCharacterLength: number = 3;

  usernameEdit: boolean = false;
  emailEdit: boolean = false;  

  username: string = 'johnny';
  password: string = '************';
  email: string = '123@gmail.com';

  usernameForm = new FormGroup({
    'username': new FormControl(this.username, [Validators.pattern('^[A-z0-9]+$')]),
  });

  emailForm = new FormGroup({
    'email': new FormControl(this.email, [Validators.email]),
  });

  passwordForm = new FormGroup({
    'password': new FormControl(this.password, [Validators.minLength(this.passwordCharacterLength)]),
  });


  editUsername(): void {

    this.usernameEdit = !this.usernameEdit;
    
  }

  editEmail(): void {

    this.emailEdit = !this.emailEdit;
  }

  editPassword(): void {

    this.callDialog();
    
  }

  back(): void {

    this.router.navigate(['users']);
  }

  callDialog(): void {

    this.dialog.open( PasswordResetComponent, { 
      height: '335px',
      width: '500px',
    });    
  }
}