import { Component, ElementRef, ViewChild } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { update } from 'cypress/types/lodash';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {

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
    
  }
}
