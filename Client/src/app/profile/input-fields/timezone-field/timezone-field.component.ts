import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-timezone-field',
  templateUrl: './timezone-field.component.html',
  styleUrls: ['./timezone-field.component.css']
})
export class TimezoneFieldComponent {
  
  editing: boolean = false;
  @Input() oldTimeZone: string = '';
  timeZoneForm: FormControl = new FormControl({value: this.oldTimeZone, disable: true}, [Validators.pattern('^[-+]{0,1}[0-9][0-9][0-9][0-9]UTC+$')])
  
  editTimeZone(): void {

    this.editing = !this.editing;       

    if(this.editing){
      this.timeZoneForm.setValue('');
      this.timeZoneForm.enable();
    }
  }

  updateTimeZone(): void {

    const newTimeZone: string = this.timeZoneForm.getRawValue();    

    if(newTimeZone !== this.oldTimeZone){
      
      // this.api.updateUserEmail(newUsername).subscribe();
      this.editTimeZone();
      this.oldTimeZone = newTimeZone;
            
    }
    else{

      this.timeZoneForm.setErrors({'duplicate': true});
    }       
  }
}
