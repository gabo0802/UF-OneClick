import { HttpErrorResponse } from '@angular/common/http';
import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { MatDialogRef } from '@angular/material/dialog';
import { Router } from '@angular/router';
import { ApiService } from 'src/app/api.service';
import { AuthService } from 'src/app/auth.service';
import { DialogsService } from 'src/app/dialogs.service';
import  passwordLength  from 'src/app/passwordLength';

@Component({
  selector: 'app-password-reset',
  templateUrl: './password-reset.component.html',
  styleUrls: ['./password-reset.component.css']
})
export class PasswordResetComponent {

  constructor(private api: ApiService, private authService: AuthService, private router: Router, private dialogs: DialogsService, private dialogRef: MatDialogRef<PasswordResetComponent>) {}

  passwordCharacterLength: number = passwordLength;

  hideOld: boolean = true;

  hideNew: boolean = true;

  passwordForm: FormGroup = new FormGroup({
    'oldPassword': new FormControl('', [Validators.required, Validators.minLength(this.passwordCharacterLength)]),
    'newPassword': new FormControl('', [Validators.required, Validators.minLength(this.passwordCharacterLength)]),
  });

  onSave(): void {    

    let oldPassword: string = this.passwordForm.get('oldPassword')?.value;
    let newPassword: string = this.passwordForm.get('newPassword')?.value;

    //checks if new password is duplicate of old password
    if(oldPassword === newPassword){
      
      this.passwordForm.get('newPassword')?.setErrors({'duplicate': true});
    }
    else{

      this.api.updateUserPassword(oldPassword, newPassword).subscribe({

        next: (res: Object) => {

          this.dialogs.successDialog("Your Password has been successfully reset. Please log back in.");

          this.dialogRef.close();
          this.authService.userLogOut();
          this.router.navigate(['login']);
        },
        error: (error: HttpErrorResponse) => {
          

          if(error.status === 409){
            this.dialogs.errorDialog("Error!", error["error"]["Error"]);
          }
          else{
            this.dialogs.errorDialog("Error!", "An unexpected error occurred please try again later.");
          }
        }        
      });
    }    
  }
}
