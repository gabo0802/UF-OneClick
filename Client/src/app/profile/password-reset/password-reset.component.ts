import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiService } from 'src/app/api.service';
import { Router } from '@angular/router';
import { AuthService } from '../../auth.service';

@Component({
  selector: 'app-password-reset',
  templateUrl: './password-reset.component.html',
  styleUrls: ['./password-reset.component.css']
})
export class PasswordResetComponent implements OnInit{
  constructor(private router: Router, private api: ApiService, private authService: AuthService) {}

  passwordCharacterLength: number = 3;

  hideOld: boolean = true;

  hideNew: boolean = true;

  passwordForm: FormGroup = {} as FormGroup;

  ngOnInit(): void {
    
    this.passwordForm = new FormGroup({
      'oldPassword': new FormControl('', [Validators.required, Validators.minLength(this.passwordCharacterLength)]),
      'newPassword': new FormControl('', [Validators.required, Validators.minLength(this.passwordCharacterLength)]),
    });
  }

  onSave(): void {
    var oldPass: string = this.passwordForm.controls['oldPassword'].value || "";
    var newPass: string = this.passwordForm.controls['newPassword'].value || "";

    if (oldPass != ""){
      if (confirm("Are you sure you want to password from " + oldPass + " to " + newPass + "?")){
        this.api.updateUserPassword(oldPass, newPass).subscribe((res) => {
          alert("Password has been changed! Logging Out...")
          this.authService.userLogOut();
          this.router.navigate(['login']);
        })
      }
    }
  }

}
