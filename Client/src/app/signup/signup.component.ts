import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiService } from '../api.service';
import { MatDialog } from '@angular/material/dialog';
import { SignupMessageComponent } from './signup-message/signup-message.component';
import { Router } from '@angular/router';

@Component({
  selector: 'app-signup',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit{

  public message: string = 'Please enter the required information below.';

  constructor(private api: ApiService, private dialog: MatDialog, private router: Router) {};

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

    
    this.api.createUser(this.signUpForm.value).subscribe( (res: Object) => {
      
      const response: string = JSON.stringify(res);
      const responseMessage = JSON.parse(response);

      //User created
      if(responseMessage["Success"] !== undefined){        
        
        this.callDialog("Success", responseMessage["Success"]);       

      }//username or email taken
      else if(responseMessage["Error"] !== undefined){
        
        this.callDialog("Error", responseMessage["Error"]);
        
      }
    });
    
  }

  callDialog(title: string, message: string){

    let dialogRef = this.dialog.open(SignupMessageComponent, {
      data: { dialogTitle: title, dialogMessage: message},      
      height: '180px',
      width: '370px',
    });

    if(title === "Success"){

      dialogRef.afterClosed().subscribe(result => {
        
        this.router.navigate(['/login']);
        this.signUpForm.reset();
      });
    }
  }

  ngOnDestroy(){

  }
}
