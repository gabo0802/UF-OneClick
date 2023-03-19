import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';

import { UsernameFieldComponent } from './username-field.component';

describe('UsernameFieldComponent', () => {
  let component: UsernameFieldComponent;
  let fixture: ComponentFixture<UsernameFieldComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        UsernameFieldComponent
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

    fixture = TestBed.createComponent(UsernameFieldComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('editing should be false initially', () => {
    expect(component.editing).toBeFalsy();
  });

  it('oldUsername should be empty before call', () => {
    expect(component.oldUsername).toEqual('');
  });

  it('Username form should be disabled initially', () => {
    expect(component.usernameForm.disabled).toBeTruthy();
  });

  it('editing variable set to true when editUsername() initially called', () =>{

    component.editUsername();

    expect(component.editing).toBeTruthy();
  });

  it('when editUsername() called form is enabled', () => {

    component.editUsername()

    expect(component.usernameForm.enabled).toBeTruthy();
  });

  it('when editUsername() called form value is empty string', () => {

    component.editUsername()

    expect(component.usernameForm.getRawValue()).toEqual('');
  });

  it('when editing if editUsername() called, editing is false and form is disabled', () => {

    //enable editing
    component.editUsername()

    //disabled editing
    component.editUsername()

    expect(component.usernameForm.disabled).toBeTruthy();
    expect(component.editing).toBeFalsy();
  });
});
