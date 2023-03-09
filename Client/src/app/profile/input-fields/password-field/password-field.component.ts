import { Component } from '@angular/core';
import { FormControl } from '@angular/forms';
import { DialogsService } from 'src/app/dialogs.service';

@Component({
  selector: 'app-password-field',
  templateUrl: './password-field.component.html',
  styleUrls: ['./password-field.component.css']
})
export class PasswordFieldComponent {

  constructor(private dialogs: DialogsService) {}

  hide: boolean = true;
  
  passwordForm: FormControl = new FormControl({value: '**********', disabled: true});  

  editPassword(): void {
    this.dialogs.passwordReset();    
  }
}
