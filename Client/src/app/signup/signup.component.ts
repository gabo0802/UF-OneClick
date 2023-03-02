import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiService } from '../api.service';
import { Router } from '@angular/router';
import { DialogsService } from '../dialogs.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit{

  public message: string = 'Please enter the required information below.';

  constructor(private api: ApiService, private dialogs: DialogsService, private router: Router) {};

  hide: boolean = true;
  passwordCharacterLength: number = 3;

  signUpForm: FormGroup = {} as FormGroup;

  ngOnInit(){

    this.signUpForm = new FormGroup({
      'username': new FormControl(null,[Validators.required, Validators.pattern('^[A-z0-9]+$')]),
      'email': new FormControl(null, [Validators.required, Validators.email]),
      'password': new FormControl(null, [Validators.required, Validators.minLength(this.passwordCharacterLength)])
    });
  }

  onSubmit(){

    type UserData = {
      username: string;
      email: string;
      password: string;
    }

    const newUser: UserData = {
      username: this.signUpForm.value.username,
      email: this.signUpForm.value.email,
      password: this.signUpForm.value.password,
    }

    this.api.createUser(newUser).subscribe(( resultMessage: string[]) => {      

      if(resultMessage[0] === "Success"){
        this.dialogs.successDialog(resultMessage[1]);
      }
      else if(resultMessage[0] === "Error"){
        this.dialogs.successDialog(resultMessage[1]);
      }
    });
    
  }

  ngOnDestroy(){

  }
}
