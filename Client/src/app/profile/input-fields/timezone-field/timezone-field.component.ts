import { HttpErrorResponse } from '@angular/common/http';
import { Component, Input } from '@angular/core';
import { FormControl, Validators } from '@angular/forms';
import { ApiService } from 'src/app/api.service';
import { AuthService } from 'src/app/auth.service';
import { DialogsService } from 'src/app/dialogs.service';

@Component({
  selector: 'app-timezone-field',
  templateUrl: './timezone-field.component.html',
  styleUrls: ['./timezone-field.component.css']
})
export class TimezoneFieldComponent {

  constructor(private api: ApiService, private dialogs: DialogsService, private authService: AuthService) {}
  
  editing: boolean = false;
  @Input() oldTimeZone: string = '';
  timeZoneForm: FormControl = new FormControl({value: this.oldTimeZone, disabled: true}, [Validators.required, Validators.pattern('^[-+]{0,1}[0-9]{4}(UTC)+$')])
  
  editTimeZone(): void {

    this.editing = !this.editing;       

    if(this.editing){
      this.timeZoneForm.setValue('UTC');
      this.timeZoneForm.enable();
    }
    else{
      this.timeZoneForm.disable();
    }
  }

  updateTimeZone(): void {

    let newTimeZone: string = this.timeZoneForm.getRawValue();    

    if(newTimeZone !== this.oldTimeZone){
      
      newTimeZone = newTimeZone.replace('UTC','');
      newTimeZone = newTimeZone.replace('+','');      

      this.api.updateTimezone(newTimeZone).subscribe({

        next: (res: Object) => {

          this.oldTimeZone = newTimeZone;
          this.editTimeZone();

        },
        error: (error: HttpErrorResponse) => {
          
          this.dialogs.errorDialog("Unexpected Error!", error.statusText + " Please try again later");                    
        }
      });
            
    }
    else{

      this.timeZoneForm.setErrors({'duplicate': true});
    }       
  }
}
