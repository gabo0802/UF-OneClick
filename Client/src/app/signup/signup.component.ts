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

  ngAfterContentChecked()	{
    this.message = this.api.getOutput()
  } *

  onSubmit(){
    
    this.api.createUser(this.signUpForm.value).subscribe( (res: Object) => {

      
      const response: string = JSON.stringify(res);
      const responseMessage = JSON.parse(response);

      if(responseMessage["Success"] !== undefined){
        console.log(responseMessage["Success"]);        
      }
      else if(responseMessage["Error"] !== undefined){
        console.log(responseMessage["Error"] );
      }
    });
    
  }

  ngOnDestroy(){

  }
}
