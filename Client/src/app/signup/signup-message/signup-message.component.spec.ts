import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SignupMessageComponent } from './signup-message.component';

describe('SignupMessageComponent', () => {
  let component: SignupMessageComponent;
  let fixture: ComponentFixture<SignupMessageComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SignupMessageComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SignupMessageComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
