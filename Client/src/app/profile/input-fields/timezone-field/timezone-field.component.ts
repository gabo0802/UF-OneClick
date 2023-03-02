import { Component, Input, OnInit } from '@angular/core';
import { FormControl, FormGroup, Validators } from '@angular/forms';

@Component({
  selector: 'app-timezone-field',
  templateUrl: './timezone-field.component.html',
  styleUrls: ['./timezone-field.component.css']
})
export class TimezoneFieldComponent implements OnInit{

  timeZoneForm: FormGroup = {} as FormGroup;
  editing: boolean = false;
  @Input() oldTimeZone: string = '';
  
  ngOnInit(): void {

    this.timeZoneForm = new FormGroup({
      'timezone': new FormControl({value: this.oldTimeZone, disable: true}, [Validators.pattern('^[-+]{0,1}[0-9][0-9][0-9][0-9]UTC+$')]),
    });
  }

  editTimeZone(): void {

    this.editing = !this.editing;
    this.timeZoneForm.enable();

    if(this.editing){
      this.timeZoneForm.get('timezone')?.setValue('');
    }
  }

  updateTimeZone(): void {

    const newTimeZone: string = this.timeZoneForm.get('timezone')?.value;    

    if(newTimeZone !== this.oldTimeZone){

      // this.api.updateUserEmail(newUsername).subscribe();
      this.editTimeZone();
      this.oldTimeZone = newTimeZone;
            
    }       
  }
}
