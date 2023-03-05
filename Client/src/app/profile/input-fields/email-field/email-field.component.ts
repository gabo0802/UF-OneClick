import { Component, Input} from '@angular/core';
import { FormControl, Validators } from '@angular/forms';

@Component({
  selector: 'app-email-field',
  templateUrl: './email-field.component.html',
  styleUrls: ['./email-field.component.css']
})
export class EmailFieldComponent {
  
  editing: boolean = false;
  @Input() oldEmail: string = '';
  emailForm: FormControl = new FormControl({value: this.oldEmail, disabled: true}, [Validators.required, Validators.email]);

  editEmail(): void {

    this.editing = !this.editing;       

    if(this.editing){
      this.emailForm.enable();
      this.emailForm.setValue('');
    }
  }

  updateEmail(): void {

    const newEmail: string = this.emailForm.getRawValue();    

    if(newEmail === this.oldEmail){
      
      this.emailForm.setErrors({'duplicate': true});
    }
    else{

      // this.api.updateUserEmail(newUsername).subscribe();
      this.editEmail();
      this.oldEmail = newEmail;
      
    }    
  }
}
