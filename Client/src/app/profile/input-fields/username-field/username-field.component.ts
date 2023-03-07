import { Component, Input, OnChanges, OnInit, SimpleChanges} from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { ApiService } from 'src/app/api.service';

@Component({
  selector: 'app-username-field',
  templateUrl: './username-field.component.html',
  styleUrls: ['./username-field.component.css']
})
export class UsernameFieldComponent {

  constructor(private api: ApiService) {}
  
  @Input() oldUsername: string = '';
  editing: boolean = false;
  usernameForm: FormControl = new FormControl({value: this.oldUsername, disabled: true}, [Validators.required, Validators.pattern('^[A-z0-9]+$')]);

  editUsername(): void {

    this.editing = !this.editing;       
    
    if(this.editing){
      this.usernameForm.enable();
      this.usernameForm.setValue('');      
    }
    else{
      this.usernameForm.disable();
    }
  }

  updateUsername(): void {
            
    const newUsername: string = this.usernameForm.getRawValue();

    if(newUsername === this.oldUsername){

      this.usernameForm.setErrors({'duplicate': true});     
      
    }
    else{

      // this.api.updateUsername(newUsername).subscribe((res) => {

      // });

      this.oldUsername = newUsername;
      this.editUsername();               
    }    
  }

}
