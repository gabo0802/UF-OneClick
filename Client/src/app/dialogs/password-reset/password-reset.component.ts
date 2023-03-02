import { Component } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-password-reset',
  templateUrl: './password-reset.component.html',
  styleUrls: ['./password-reset.component.css']
})
export class PasswordResetComponent {

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
    
  }

}
