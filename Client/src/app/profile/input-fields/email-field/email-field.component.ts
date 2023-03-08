import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input} from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ApiService } from 'src/app/api.service';
import { AuthService } from 'src/app/auth.service';
import { DialogsService } from 'src/app/dialogs.service';

@Component({
  selector: 'app-email-field',
  templateUrl: './email-field.component.html',
  styleUrls: ['./email-field.component.css']
})
export class EmailFieldComponent {

  constructor(private api: ApiService, private dialogs: DialogsService, private router: Router, private authService: AuthService) {}
  
  editing: boolean = false;
  @Input() oldEmail: string = '';
  emailForm: FormControl = new FormControl({value: this.oldEmail, disabled: true}, [Validators.required, Validators.email]);

  editEmail(): void {

    this.editing = !this.editing;       

    if(this.editing){
      this.emailForm.enable();
      this.emailForm.setValue('');
    }
    else{
      this.emailForm.disable();
    }
  }

  updateEmail(): void {

    const newEmail: string = this.emailForm.getRawValue();    

    if(newEmail === this.oldEmail){
      
      this.emailForm.setErrors({'duplicate': true});
    }
    else{

      this.api.updateUserEmail(newEmail).subscribe({

        next: (res: Object) => {

          this.oldEmail = newEmail;
          this.editEmail();
          this.dialogs.successDialog("Please verify your email before logging back in.");
          this.authService.userLogOut();           
          this.router.navigate(['login']);
        },
        error: (error: HttpErrorResponse) => {

          if(error.status == 409){
            
            this.emailForm.setErrors({'taken': true});
          }
          else{
            this.dialogs.errorDialog("Unexpected Error!", error.statusText + " Please try again later");
          }          
        }
      });
      
    }    
  }
}
