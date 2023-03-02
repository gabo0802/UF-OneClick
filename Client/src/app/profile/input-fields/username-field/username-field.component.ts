import { Component, Input} from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';
import { ApiService } from 'src/app/api.service';

@Component({
  selector: 'app-username-field',
  templateUrl: './username-field.component.html',
  styleUrls: ['./username-field.component.css']
})
export class UsernameFieldComponent {

  constructor(private api: ApiService) {}

  usernameForm: FormGroup = {} as FormGroup;
  @Input() oldUsername: string = '';
  editing: boolean = false;  

  ngOnInit(): void {

    this.usernameForm = new FormGroup({
      'username': new FormControl({value: '', disabled: true}, [Validators.required, Validators.pattern('^[A-z0-9]+$')]),
    });
  }

  editUsername(): void {

    this.editing = !this.editing;   
    this.usernameForm.enable();
    
    if(this.editing){
      this.usernameForm.get('username')?.setValue('');
    }
  }

  updateUsername(): void {
    
    
    const newUsername: string = this.usernameForm.get('username')?.value;

    if(newUsername === this.oldUsername){
          
      this.usernameForm.get('username')?.setErrors({'duplicate': true});
    }
    else{

      this.api.updateUsername(newUsername).subscribe((res) => {

      });
      
          
    } 
  }

}
