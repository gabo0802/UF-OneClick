import { HttpClientModule } from '@angular/common/http';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { MaterialDesignModule } from 'src/app/material-design/material-design.module';
import { MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog';

import { PasswordResetComponent } from './password-reset.component';

describe('PasswordResetComponent', () => {
  let component: PasswordResetComponent;
  let fixture: ComponentFixture<PasswordResetComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ 
        PasswordResetComponent
       ],
       imports: [
        MaterialDesignModule, 
        HttpClientModule,
        FormsModule,
        ReactiveFormsModule,
        BrowserAnimationsModule,
       ],
       providers: [
        { provide: MAT_DIALOG_DATA, useValue: {title: null, message: null}},
        { provide: MatDialogRef, useValue: {title: null}}        
       ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(PasswordResetComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });

  it('oldPassword and newPassword input fields hidden by default', ()=> {

    expect(component.hideOld).toBeTruthy();
    expect(component.hideNew).toBeTruthy();
  });

  it('OldPassword and NewPassword form values should be empty string intially', () => {

    expect(component.passwordForm.get('newPassword')?.value).toEqual('');
    expect(component.passwordForm.get('oldPassword')?.value).toEqual('');
  });

  it('if newPassword is idenitical to oldPassword throws form duplicate error', ()=> {

    component.passwordForm.get('newPassword')?.setValue('Ja1@4');
    component.passwordForm.get('oldPassword')?.setValue('Ja1@4');

    component.onSave();

    expect(component.passwordForm.get('newPassword')?.hasError('duplicate'));
  });
});
