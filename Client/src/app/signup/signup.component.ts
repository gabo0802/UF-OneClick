import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit{

  public message: string = 'Please enter the required information below.';

  constructor(private api: ApiService) {};

  hide = true;

  signUpForm: FormGroup = {} as FormGroup;

  ngOnInit(){

    this.signUpForm = new FormGroup({
      'username': new FormControl(null,[Validators.required, Validators.pattern('[A-z0-9]+')]),
      'email': new FormControl(null, [Validators.required, Validators.email]),
      'password': new FormControl(null, [Validators.required, Validators.minLength(3)])
    });
  }

  onSubmit(){
    
    this.message = this.api.getOutput()

    this.api.createUser(this.signUpForm.value)

  }
}
