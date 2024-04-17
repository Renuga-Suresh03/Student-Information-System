// JavaScript for student pages

// Function to handle student login form submission
function handleStudentLogin() {
    var regNo = document.getElementById('regNo').value;
    var dob = document.getElementById('dob').value;

    // Perform client-side validation if needed

    // Send login request to backend API
    fetch('/api/student/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            reg_no: regNo,
            dob: dob
        }),
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Failed to login');
        }
        return response.json();
    })
    .then(data => {
        // Redirect to student home page or display success message
        console.log('Student login successful:', data);
        window.location.href = '/student/home.html';
    })
    .catch(error => {
        // Display error message to user
        console.error('Student login failed:', error);
        alert('Student login failed. Please try again.');
    });
}

// Function to handle navigation to student home page
function navigateToStudentHome() {
    window.location.href = '/student/home.html';
}

// Function to handle navigation to student profile page
function navigateToStudentProfile() {
    window.location.href = '/student/profile.html';
}

// Function to handle navigation to student marks page
function navigateToStudentMarks() {
    window.location.href = '/student/marks.html';
}

// Function to handle navigation to student attendance page
function navigateToStudentAttendance() {
    window.location.href = '/student/attendance.html';
}

// Function to display assessment marks for a student
function displayAssessmentMarks(examType) {
    // Implement functionality to fetch and display assessment marks
    console.log('Displaying assessment marks for exam:', examType);
}

// Function to display attendance records for a student
function displayAttendanceRecords() {
    // Implement functionality to fetch and display attendance records
    console.log('Displaying attendance records');
}
