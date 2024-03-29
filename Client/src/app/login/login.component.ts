import { HttpErrorResponse } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ApiService } from '../api.service';
import { AuthService } from '../auth.service';
import { DialogsService } from '../dialogs.service';


@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  loginForm: FormGroup = {} as FormGroup;
  hide: boolean = true;

  constructor(private api: ApiService, private router: Router, private dialogs: DialogsService, private authService: AuthService) {};  

  ngOnInit(){
    this.loginForm = new FormGroup({
      'username': new FormControl(null, [Validators.required, Validators.pattern('^[A-z0-9]+$')]),
      'password': new FormControl(null, Validators.required)
    });
  }

  onSubmit(){

    this.api.login(this.loginForm.value).subscribe({

      next: (res) => {

        this.authService.userLogIn();

        if(this.authService.isAdmin()){
          this.router.navigate(['admin'])
        }
        else{
          this.router.navigate(['users']);
        }        
        this.loginForm.reset();
      },
      error: (error: HttpErrorResponse) => {
        
        if(error.status === 504){
          this.dialogs.errorDialog("Error Logging In!", error.statusText + " Please try again later.");
        }
        else{
          this.dialogs.errorDialog("Error Logging In!", error["error"]["Error"]);
        }     
        
      }
    });     
  }

}
