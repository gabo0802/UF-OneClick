import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-email-field',
  templateUrl: './email-field.component.html',
  styleUrls: ['./email-field.component.css']
})
export class EmailFieldComponent implements OnInit{

  emailForm: FormGroup = {} as FormGroup;
  editing: boolean = false;
  @Input() oldEmail: string = '';
  

  ngOnInit(): void {
    this.emailForm = new FormGroup({
      'email': new FormControl({value: this.oldEmail, disabled: true}, [Validators.required, Validators.email]),
    });
  }

  editEmail(): void {

    this.editing = !this.editing;
    this.emailForm.enable();

    if(this.editing){
      this.emailForm.get('email')?.setValue('');
    }
  }

  updateEmail(): void {

    const newEmail: string = this.emailForm.get('email')?.value;    

    if(newEmail === this.oldEmail){
      
      this.emailForm.get('email')?.setErrors({'duplicate': true});
    }
    else{

      // this.api.updateUserEmail(newUsername).subscribe();
      this.editEmail();
      this.oldEmail = newEmail;
      
    }    
  }
}
