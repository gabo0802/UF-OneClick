import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialog } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { ApiService } from '../api.service';
import { AuthService } from '../auth.service';
import { LoginMessageComponent } from './login-message/login-message.component';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {

  loginForm: FormGroup = {} as FormGroup;
  hide: boolean = true;

  constructor(private api: ApiService, private dialog: MatDialog, private router: Router, private authService: AuthService) {};  

  ngOnInit(){
    this.loginForm = new FormGroup({
      'username': new FormControl(null, [Validators.required, Validators.pattern('^[A-z0-9]+$')]),
      'password': new FormControl(null, Validators.required)
    });
  }

  onSubmit(){

    this.api.login(this.loginForm.value).subscribe( (resultMessage: string[]) => {

      if(resultMessage[0] === "Success"){
        
        this.authService.userLogIn();
        this.router.navigate(['users']);
        this.loginForm.reset();
      }
      else{
        this.callDialog(resultMessage[0], resultMessage[1]);
      }
    })

    
  }

  callDialog(title: string, message: string){

    this.dialog.open(LoginMessageComponent, {
      data: { dialogTitle: title, dialogMessage: message},      
      height: '180px',
      width: '370px',
    });    
  }

}
