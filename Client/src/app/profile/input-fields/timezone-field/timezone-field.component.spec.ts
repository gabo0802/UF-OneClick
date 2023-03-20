import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';

import { TimezoneFieldComponent } from './timezone-field.component';

describe('TimezoneFieldComponent', () => {
  let component: TimezoneFieldComponent;
  let fixture: ComponentFixture<TimezoneFieldComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        TimezoneFieldComponent
       ],
       imports: [
        MaterialDesignModule, 
        HttpClientModule,
        FormsModule,
        ReactiveFormsModule,
        BrowserAnimationsModule
       ],

    })
    .compileComponents();

    fixture = TestBed.createComponent(TimezoneFieldComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('editing should be false initially', () => {
    expect(component.editing).toBeFalsy();
  });

  it('oldTimeZone should be empty before call', () => {
    expect(component.oldTimeZone).toEqual('');
  });

  it('Time Zone form should be disabled initially', () => {
    expect(component.timeZoneForm.disabled).toBeTruthy();
  });
});
