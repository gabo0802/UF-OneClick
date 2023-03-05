import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ApiService } from 'src/app/api.service';
import { AuthService } from 'src/app/auth.service';
import { DialogsService } from 'src/app/dialogs.service';

@Component({
  selector: 'app-password-reset',
  templateUrl: './password-reset.component.html',
  styleUrls: ['./password-reset.component.css']
})
export class PasswordResetComponent {

  constructor(private api: ApiService, private authService: AuthService, private router: Router, private dialogs: DialogsService) {}

  passwordCharacterLength: number = 3;

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

      this.passwordForm.setErrors({'duplicate': true});
    }
    else{

      this.api.updateUserPassword(oldPassword, newPassword).subscribe( (res) => {

        this.dialogs.successDialog("Your Password has been successfully reset. Please log back in.");

        this.authService.userLogOut();
        this.router.navigate(['login']);
      });
    }    
  }

}
