import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  loginForm: FormGroup = {} as FormGroup;
  constructor(private api: ApiService) {};
  hide: boolean = true;

  ngOnInit(){
    this.loginForm = new FormGroup({
      'username': new FormControl(null, null),
      'password': new FormControl(null, null)
    });
  }

  onSubmit(){
    this.api.Login(this.loginForm.value)
    this.loginForm.reset();
  }

}
