import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';

import { PasswordFieldComponent } from './password-field.component';

describe('PasswordFieldComponent', () => {
  let component: PasswordFieldComponent;
  let fixture: ComponentFixture<PasswordFieldComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        PasswordFieldComponent
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

    fixture = TestBed.createComponent(PasswordFieldComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('password input field should be hidden by default', () => {
    
    expect(component.hide).toBeTruthy();
  });

  it('form should be disabled', () => {

    expect(component.passwordForm.disabled).toBeTruthy();
  });

  it('form value should be asterisks', () => {

    expect(component.passwordForm.getRawValue()).toEqual('**********');
  })
});
