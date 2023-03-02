import { Component, OnInit } from '@angular/core';
import { FormControl, FormGroup } from '@angular/forms';
import { DialogsService } from 'src/app/dialogs.service';

@Component({
  selector: 'app-password-field',
  templateUrl: './password-field.component.html',
  styleUrls: ['./password-field.component.css']
})
export class PasswordFieldComponent implements OnInit{

  constructor(private dialogs: DialogsService) {}

  hide: boolean = true;
  passwordCharacterLength: number = 3; 
  passwordForm: FormGroup = {} as FormGroup;

  ngOnInit(): void {
    
    this.passwordForm = new FormGroup({
      'password': new FormControl('************'),
    });
  }

  editPassword(): void {
    this.dialogs.passwordReset();    
  }
}
