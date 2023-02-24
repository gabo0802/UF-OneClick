import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent {

  hide: boolean = true;
  passwordCharacterLength: number = 3;

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

}
